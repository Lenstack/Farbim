import {useEffect, useState} from "react";
import {BaseUrl} from "@/services";

export const useFetch = (url: string, options: object) => {
    const [data, setData] = useState([]);
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        setLoading(true);
        fetch(BaseUrl + url, options)
            .then((response) => response.json())
            .then((data) => setData(data))
            .catch((err) => setError("Error: " + err))
            .finally(() => setLoading(false));
    }, [url]);

    return {data, error, loading};
};