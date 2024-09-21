import { api } from "@/lib/api";
import BudgetRequest from "@/models/budget-request";

interface FetchBudgetItemsResponse {
  data: BudgetRequest[];
}

export const fetchBudgetItems = async (): Promise<BudgetRequest[]> => {
  const response = await api.get<FetchBudgetItemsResponse>("/items");
  const { data } = response.data;
  return data;
};

interface CreateBudgetItemRequest {
    title: string;
    quantity: number;
    price: number;
  }
  
  interface CreateBudgetItemResponse {
    data: BudgetRequest;
  }
  
  export const createBudgetItem = async (body: CreateBudgetItemRequest) => {
    const response = await api.post<CreateBudgetItemResponse>("/items", body);
    const { data } = response.data;
    return data;
  };

  // New function to fetch an item by its ID
interface FetchBudgetItemByIdResponse {
    data: BudgetRequest;
  }
  
  export const fetchBudgetItemById = async (id: number): Promise<BudgetRequest> => {
    const response = await api.get<FetchBudgetItemByIdResponse>(`/items/${id}`);
    const { data } = response.data;
    return data;
  };