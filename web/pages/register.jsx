import Link from "next/link";

const Register = () => {
  const registerHandler = (e) => {
    e.preventDefault();
  };
  return (
    <>
      <div className="container">
        <form onSubmit={(e) => registerHandler(e)}>
          <input type="text" placeholder="Username" />
          <input type="text" placeholder="Fullname" />
          <button type="submit">Register</button>
          <span>
            Have an account ?{" "}
            <Link href="/login">
              <a>Login</a>
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

export default Register;
