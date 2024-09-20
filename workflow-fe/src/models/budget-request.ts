interface budgetRequest {
    id: number;
    title: string;
    amount: number;
    quantity: string;
    status: "PENDING" | "APPROVED" | "REJECTED";
    owner_id: number;
};
export default budgetRequest