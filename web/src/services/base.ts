const baseUrl = "http://localhost:8080/v1/"

export const FetchingApi = async (moduleApi: string, optionsApi: object) => {
    return await fetch(baseUrl + moduleApi, optionsApi)
        .then(response => response.json())
        .then(jsonData => {
            return jsonData
        })
        .catch(error => {
            return error
        })
}