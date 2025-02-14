import { type MetaFunction } from "@remix-run/node";
import { useState } from "react";

export const meta: MetaFunction = () => {
  return [
    { title: "URL Encode and Decode" },
    { name: "description", content: "Easily Encode and Decode URL with our online tool." },
  ];
};

export default function Index() {
  const [input, setInput] = useState("");
  const [result, setResult] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [encodeOnlyParams, setEncodeOnlyParams] = useState(false);
  const [isCopied, setIsCopied] = useState(false);

  const handleApiCall = async (operation: "encode" | "decode") => {
    if (!input.trim()) {
      setError("Please enter a URL");
      return;
    }

    setLoading(true);
    setError("");

    try {
      const response = await fetch(
        `/${ operation }?encode_only_params=${ encodeOnlyParams }`,
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

  const handleCopy = async () => {
    try {
      await navigator.clipboard.writeText(result);
      setIsCopied(true);
      setTimeout(() => setIsCopied(false), 2000);
    } catch (err) {
      console.error("Failed to copy text: ", err);
    }
  };

  return (
    <div className="flex flex-col justify-center items-center h-screen w-screen">
      <main className="flex flex-col items-center justify-center rounded-3xl gap-6 w-8/12">
        <h1 className="leading text-6xl font-bold text-gray-800 dark:text-gray-100">
          Paste your link below to encode or decode it
        </h1>
        <textarea
          placeholder="https://url-dencoder.vercel.app/?title=hello&description=world"
          className="px-2 py-1 w-full border-b border-gray-800 text-lg min-h-24 resize-none"
          autoFocus
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
            disabled={loading || !input}
            className="border border-gray-800 px-3 py-1 disabled:opacity-50"
          >
            {loading ? "Encoding..." : "Encode"}
          </button>
          <button
            type="button"
            onClick={() => handleApiCall("decode")}
            disabled={loading || !input}
            className="border border-gray-800 px-3 py-1 disabled:opacity-50"
          >
            {loading ? "Decoding..." : "Decode"}
          </button>
        </div>
        <div className="mt-3 w-full h-[160px]">
          {result && (
            <>
              <div className="flex justify-between mb-2">
                <h2 className="text-xl font-semibold ">Result:</h2>
                <button
                  onClick={handleCopy}
                  className="bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-4 rounded-lg shadow-md transition duration-300"
                >
                {isCopied ? (
                    <span>Copied</span>
                  ) : <span> Copy to Clipboard</span>}
                </button>
              </div>
              <div className="p-4 bg-gray-100 dark:bg-gray-800 rounded break-words">
                {result}
              </div>
            </>
          )}
          {error && (
            <>
              <h2 className="text-xl font-semibold mb-2">Error:</h2>
              <div className="p-4 bg-gray-100">
                <span className="text-red-500 dark:text-red-400">{error}</span>
              </div>
            </>
          )}
        </div>
      </main>
    </div>
  );
}