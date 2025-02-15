export const onRequest = async ({ request }: { request: Request }) => {
  const url = new URL(request.url);

  if (url.pathname === "/encode") {
    return new Response("Encoding API");
  }

  if (url.pathname === "/decode") {
    return new Response("Decoding API");
  }

  return new Response("Not Found", { status: 404 });
};