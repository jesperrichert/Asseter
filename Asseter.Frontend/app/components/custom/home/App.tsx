import { useContext } from "react";
import { AuthContext } from "~/context/auth.context";
import Cookies from "js-cookie";

export function App() {
  const userSession = useContext(AuthContext);

  const logout = () => {
    Cookies.remove("session");
    window.open("/auth", "_self");
  };

  return (
    <div>
      {userSession?.user?.username ?? "UNKNOWN"}
      <div>
        <button onClick={() => logout()} className="p-2 border-2">
          Logout
        </button>
      </div>
    </div>
  );
}
