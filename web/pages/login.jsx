import React from "react";
import Link from "next/link";
import { useRouter } from "next/router";
import axios from "axios";

const Login = () => {
  const [username, setUsername] = React.useState("");
  const [loading, setLoading] = React.useState(false);
  const router = useRouter();

  const loginHandler = (e) => {
    e.preventDefault();
    if (loading) {
      return;
    }
    setLoading(true);
    if (username === "") {
      alert("Username can't blank");
    }
    axios
      .post("http://54.169.75.182:8080/login", {
        username,
      })
      .then((res) => {
        alert(res.data.message);
        if (res.data.data) {
          if (typeof window !== "undefined") {
            window.localStorage.setItem("USER", JSON.stringify(res.data.data));
            router.replace("/");
          }
        }
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };
  return (
    <>
      <div className="container">
        <form onSubmit={(e) => loginHandler(e)}>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => {
              // Check id valid
              setUsername(e.target.value);
            }}
          />
          <button type="submit">{loading ? "Loading" : "Login"}</button>
          <span>
            Doesn't have an account ?{" "}
            <Link href="/register">
              <a>Register</a>
            </Link>
          </span>
        </form>
      </div>
      <style jsx>{`
        .container {
          width: 100%;
          height: 100vh;
          overflow: hidden;
          display: flex;
          align-items: center;
          justify-content: center;
        }
        .container form {
          display: flex;
          flex-direction: column;
        }
        .container form input,
        .container form button {
          padding: 16px 32px;
          margin-bottom: 16px;
        }
        .container form input {
          border: 1px solid rgba(100, 100, 100, 0.5);
          font-size: 2em;
        }
        .container form input:focus {
          outline: none;
        }
        .container form button {
          border: none;
          background-color: #000;
          color: #fff;
          font-weight: bold;
          cursor: pointer;
          transition: all 0.25s ease;
        }
        .container form button:hover {
          background-color: #333;
        }
      `}</style>
    </>
  );
};

export default Login;
