export const getLocalStorage = (key: string) => {
    return localStorage.getItem(key)
}

export const setLocalStorage = (key: string, value: string) => {
    return localStorage.setItem(key, value)
}

export const clearItemLocalStorage = (key: string) => {
    return localStorage.removeItem(key)
}

export const clearLocalStorage = () => {
    return localStorage.clear()
}