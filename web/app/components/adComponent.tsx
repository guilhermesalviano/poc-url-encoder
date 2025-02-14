import { useEffect } from "react";

export default function AdComponent() {
    useEffect(() => {
        const script = document.createElement("script");
        script.src =
            "https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-2721142105042988";
        script.async = true;
        script.crossOrigin = "anonymous";
        document.head.appendChild(script);

        return () => {
            document.head.removeChild(script);
        };
    }, []);

    return (
        <>
            <h1>Página com Anúncios</h1>
            <ins
                className="adsbygoogle"
                style={{ display: "block" }}
                data-ad-client="ca-pub-2721142105042988"
                data-ad-slot="1234567890"
                data-ad-format="auto"
                data-full-width-responsive="true"
            />
            <script>
                {`(adsbygoogle = window.adsbygoogle || []).push({});`}
            </script>
        </>
    );
}