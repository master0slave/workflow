"use client";
import React, { useState, useEffect } from 'react';
import BudgetRequestDataTable from "@/components/BudgetRequestDataTable";
import { fetchBudgetItems } from '@/service/budget-item';
import BudgetRequest  from '@/models/budget-request';
import Header from '@/components/Header';

const Home = () => {
  // Explicitly type the state variable
  const [items, setItems] = useState<BudgetRequest[]>([]);

  useEffect(() => {
    const loadItems = async () => {
      try {
        const fetchedItems = await fetchBudgetItems();
        setItems(fetchedItems);
      } catch (error) {
        console.error('Failed to fetch items:', error);
        // Handle errors appropriately
      }
    };

    loadItems();
  }, []);

  return (
    <div>
      <Header />
      <main className="container mx-auto mt-4">
        <BudgetRequestDataTable items={items} />
      </main>
    </div>
  );
};

export default Home;
