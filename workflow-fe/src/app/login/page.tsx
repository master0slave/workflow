"use client";

import { api } from "@/lib/api";
import axios from "axios";
import { useRouter } from "next/navigation";
import { FormEvent, useState } from "react";

function LoginPage() {
  const [credential, setCredential] = useState({
    username: "",
    password: "",
  });
  const router = useRouter();

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    try {
      await api.post("/login", {
        username: credential.username,
        password: credential.password,
      });
      alert("success");
      router.push("/");
    } catch (err) {
      if (axios.isAxiosError<{ message: string }>(err)) {
        alert(err.response?.data.message);
      }
    }
  };
  return (
    <form className="border rounded p-4 shadow-md max-w-sm m-auto mt-10 bg-white" onSubmit={handleSubmit}>
    <h1 className="text-xl font-bold text-center mb-6">Please sign in</h1>
    <div className="mb-4">
      <label htmlFor="username" className="block text-sm font-medium text-gray-700">Email address</label>
      <input
        type="text"
        name="username"
        id="username"
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        placeholder="admin2@test.com"
        required
        onChange={(event) =>
          setCredential({
            ...credential,
            username: event.target.value,
          })
        }
      />
    </div>
    <div className="mb-6">
      <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
      <input
        type="password"
        name="password"
        id="password"
        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
        placeholder="Enter your password"
        required
        onChange={(event) =>
          setCredential({
            ...credential,
            password: event.target.value,
          })
        }
      />
    </div>
    <button type="submit" className="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">
      Sign in
    </button>
  </form>
  )

}

export default LoginPage;