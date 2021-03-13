import React from "react";
import Link from "next/link";
import { useRouter } from "next/router";

export default function Navbar() {
  const [user, setUser] = React.useState(undefined);
  const router = useRouter();
  React.useEffect(() => {
    // Check client of Server
    if (window !== undefined) {
      if (window.localStorage.getItem("USER") !== undefined) {
        setUser(JSON.parse(window.localStorage.getItem("USER")));
        // "{"username": "argana"}"
        // user.username
        /*
        {
          "username": "argana"
        }
        var.username
        */
      }
    }
  }, []);
  return (
    <>
      <nav>
        {!user ? (
          <>
            <div></div>
            <div>
              <Link href="/login">
                <a>Login</a>
              </Link>
              <Link href="/register">
                <a>Register</a>
              </Link>
            </div>
          </>
        ) : (
          <>
            <span>{user.fullname}</span>
            <div>
              <Link href="/keranjang">
                <a>Keranjang</a>
              </Link>

              <span
                onClick={(e) => {
                  if (window !== undefined) {
                    if (window.localStorage.getItem("USER") !== undefined) {
                      window.localStorage.removeItem("USER");
                      router.replace("/login");
                    }
                  }
                }}
              >
                Logout
              </span>
            </div>
          </>
        )}
      </nav>
      <style jsx>
        {`
          nav {
            padding: 24px 32px;
            display: flex;
            justify-content: space-between;
            align-items: center;
          }
          nav span {
            background-color: rgba(200, 200, 200, 0.2);
            padding: 16px 32px;
            border-radius: 6px;
            transition: all 0.25s ease;
            cursor: pointer;
          }
          nav span:hover {
            background-color: rgba(0, 0, 0, 0.8);
            color: #fff;
          }
          nav a {
            background-color: rgba(200, 200, 200, 0.2);
            padding: 16px 32px;
            border-radius: 6px;
            transition: all 0.25s ease;
          }
          nav a:not(:last-child) {
            margin-right: 16px;
          }
          nav a:hover {
            background-color: rgba(0, 0, 0, 0.8);
            color: #fff;
          }
        `}
      </style>
    </>
  );
}
