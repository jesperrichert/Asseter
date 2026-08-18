import { createContext, useEffect, useState } from "react";

export type AuthToken = string|null

export const AuthContext = createContext<AuthToken>(null)

export function Auth({ children }: { children: React.ReactNode }) {
    const [auth, setAuth] = useState<AuthToken|null>(null)

    useEffect(() => {
        if (window.location.href.split("/").includes("auth")) return
        try {
            const cookie = document.cookie.split("session=")[1].split(";")[0]
            setAuth(cookie)
        } catch (e) {
            window.open("/auth", "_self")
        }
        // TODO: Check Token... With /api/auth/@me
    }, [])

    return <AuthContext value={auth}>{children}</AuthContext>
}