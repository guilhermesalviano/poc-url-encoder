export function processContent(
  operation: "encode" | "decode",
  content: string,
  encodeOnlyParams: boolean
): string {
  try {
    if (operation === "encode") {
      if (encodeOnlyParams) {
        const [url, params] = content.split("?")
        return `${url}?${encodeURIComponent(params)}`
      }
        return encodeURIComponent(content)
    } else {
      return decodeURIComponent(content)
    }
    
  } catch (error) {
    throw new Error(
        "URL inválida para processamento de parâmetros"
    );
  }
}