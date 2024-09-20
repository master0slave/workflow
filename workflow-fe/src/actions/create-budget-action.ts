"use server";


import { createBudgetItem } from "@/service/budget-item";
import axios from "axios";
// import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";

interface CreateErrorResponse {
  message: InvalidField[];
}

interface InvalidField {
  Field: string;
  Reason: string;
}

export async function createBudgetAction(formData: FormData) {
  const rawFormData = {
    title: String(formData.get("title")), // get from <input name="title" />
    price: Number(formData.get("price")), // get from <input name="price" />
    quantity: 1,
  };
  let redirectPath = "/";

  try {
    await createBudgetItem(rawFormData);
  } catch (err: unknown) {
    if (axios.isAxiosError<CreateErrorResponse>(err)) {
      const messages = err.response?.data.message;

      const errors: { [key: string]: string } = {};
      messages?.forEach((message) => {
        const field = message.Field.toLowerCase();
        const reason = message.Reason;
        errors[field] = reason;
      });

      const params = new URLSearchParams();
      params.set("errors", JSON.stringify(errors));

      redirectPath = "/?" + params.toString();
    }
  } finally {
    // should not call redirect inside try/catch block
    // so I decided to place it in finally block instead
    redirect(redirectPath); 
  }
}