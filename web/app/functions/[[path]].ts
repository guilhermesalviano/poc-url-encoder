export const onRequest = async ({ request }: { request: Request }) => {
  return new Response("Hello from Cloudflare!", { status: 200 });
};
