import { API_BASE_URL } from "../config/api";
export const getCauseImage = (url, fallback) => { 
    
    // const BASE_URL = "http://localhost:8080";

    const BASE_URL = API_BASE_URL;
  
    if (!url || url.trim() === "") return fallback;
  
    // if (url.startsWith("http")) return url;
    // Change fallback to url fo http 
    if (url.startsWith("http")) return fallback;

  
    return `${BASE_URL}${url}`;
  };