export const onRequest = async ({ request }: { request: Request }) => {
  return new Response("Hello from the index function!", { status: 200 });
};
