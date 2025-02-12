import { useEffect } from "react";
import { useLocation } from "@remix-run/react";

export default function AdComponent() {
    const location = useLocation();

    useEffect(() => {
        handleAds();
    }, [location.key]);

    const handleAds = () => {
        if (!document.getElementById("adsbygoogleaftermount")) {
            const script = document.createElement("script");
            script.id = "adsbygoogleaftermount";
            script.type = "text/javascript";
            script.async = true;
            script.src =
                "https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-2721142105042988";
            script.crossOrigin = "anonymous";
            document.head.appendChild(script);
        }
    };

    return (
        <div className="mt-4">
            <div key={location.key}>
                <ins
                    className="adsbygoogle"
                    style={{ display: "block" }}
                    data-ad-client="ca-pub-2721142105042988"
                    data-ad-slot="xxxxxxxxxx" // Replace with your actual ad slot ID
                    data-ad-format="auto"
                    data-full-width-responsive="true"
                ></ins>
            </div>
        </div>
    );
}
