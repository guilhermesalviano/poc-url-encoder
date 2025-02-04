import type { MetaFunction } from "@remix-run/node";
import { useState } from "react";

export const meta: MetaFunction = () => {
  return [
    { title: "URL Encode/Decode" },
    { name: "description", content: "Quickly and effortlessly encode or decode URLs..." },
  ];
};

export default function Index() {
  const [input, setInput] = useState("");
  const [result, setResult] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [encodeOnlyParams, setEncodeOnlyParams] = useState(true);

  const handleApiCall = async (operation: "encode" | "decode") => {
    if (!input.trim()) {
      setError("Please enter a URL");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const response = await fetch(
        `http://localhost:3333/${operation}?encode_only_params=${encodeOnlyParams}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ content: input }),
        }
      );

      if (!response.ok) throw new Error("API request failed");

      const data = await response.json();
      setResult(data.content || "");
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to process URL");
      setResult("");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex flex-col justify-center items-center h-screen w-screen">
      <header className="flex flex-col absolute top-8 left-12">
        <div className="w-[134px]">
          <img
            src="/logo-light.png"
            alt="guilherme"
            className="block w-full dark:hidden pl-5 object-fill"
          />
        </div>
      </header>

      <main className="flex flex-col items-center justify-center rounded-3xl gap-6 w-8/12">
        <h1 className="leading text-6xl font-bold text-gray-800 dark:text-gray-100">
          Paste your link below to encode or decode it
        </h1>

        <input
          type="text"
          placeholder="https://test.com?msg=olá"
          className="px-4 py-2 w-full border-b border-gray-800 text-lg"
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />

        <div className="flex items-center gap-4">
          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              checked={encodeOnlyParams}
              onChange={(e) => setEncodeOnlyParams(e.target.checked)}
              className="w-4 h-4"
            />
            Encode only parameters
          </label>
        </div>

        <div className="flex content-end gap-4">
          <button
            type="button"
            onClick={() => handleApiCall("encode")}
            disabled={loading}
            className="border border-gray-800 px-3 py-1 disabled:opacity-50"
          >
            {loading ? "Encoding..." : "Encode"}
          </button>
          <button
            type="button"
            onClick={() => handleApiCall("decode")}
            disabled={loading}
            className="border border-gray-800 px-3 py-1 disabled:opacity-50"
          >
            {loading ? "Decoding..." : "Decode"}
          </button>
        </div>

        {result && (
          <div className="mt-4 w-full">
            <h2 className="text-xl font-semibold mb-2">Result:</h2>
            <div className="p-4 bg-gray-100 dark:bg-gray-800 rounded break-words">
              {result}
            </div>
          </div>
        )}

        {error && (
          <div className="mt-4 text-red-500 dark:text-red-400">
            Error: {error}
          </div>
        )}
      </main>
    </div>
  );
}