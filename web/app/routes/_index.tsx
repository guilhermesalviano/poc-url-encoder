import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "URL Encode/Decode" },
    { name: "description", content: "Quickly and effortlessly encode or decode URLs with our user-friendly online tool. Perfect for developers, data professionals, and anyone working with web applications, this tool ensures your URLs are web-safe and compliant with standards." },
  ];
};

export default function Index() {
  return (
    <div className="flex flex-col justify-center items-center h-screen w-screen">
      <header className="flex flex-col absolute top-8 left-12">
        <div className="w-[134px]"> {/*border border-gray-800 w-14 h-14 z-0"*/}
          <img
            src="/logo-light.png"
            alt="guilherme"
            className="block w-full dark:hidden pl-5 object-fill"
          />
        </div>
      </header>
      
      <main className="flex flex-col items-center justify-center rounded-3xl gap-6 w-8/12">
        <h1 className="leading text-6xl font-bold text-gray-800 dark:text-gray-100">
          paste the link below that you
          want to encode or decode
        </h1>
        <input type="text" placeholder="https://test.com?msg=olá" className="px-4 py-2 w-full border-b border-gray-800 text-lg" />
        <div className="flex content-end gap-4">
          <button type="button" className="border border-gray-800 px-3 py-1">encode</button>
          <button type="button" className="border border-gray-800 px-3 py-1">decode</button>
        </div>
      </main>
    </div>
  );
}
