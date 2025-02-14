import type { ActionFunctionArgs } from "@remix-run/node";
import { processContent } from "~/utils/processContent";

export async function action({ request }: ActionFunctionArgs) {
  try {
    const url = new URL(request.url);
    const encodeOnlyParams = url.searchParams.get("encode_only_params") === "true";
    
    const { content } = await request.json();
    
    if (!content) {
      return Response.json(
        { error: "Content is required" },
        { status: 400 }
      );
    }

    const operation = url.pathname.endsWith("/encode") ? "encode" : "decode";
    const processed = processContent(operation, content, encodeOnlyParams);

    return Response.json({ content: processed });
  } catch (error) {
    return Response.json(
      { error: "Failed to process request" },
      { status: 500 }
    );
  }
}

export function headers() {
  return {
    "Access-Control-Allow-Origin": "http://localhost:5173",
    "Access-Control-Allow-Methods": "POST",
    "Access-Control-Allow-Headers": "Content-Type",
  };
}