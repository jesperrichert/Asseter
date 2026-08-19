import { useContext } from "react"
import { AuthContext } from "~/context/auth.context"

export async function clientLoader() {
  // ignore
}
clientLoader.hydrate = false;

export function App() {
    const userSession = useContext(AuthContext)
    return <div>
        {userSession?.user?.username ?? "UNKNOWN"}
    </div>
}