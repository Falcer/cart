import React from "react";
import Link from "next/link";

export default function Navbar() {
  return (
    <>
      <nav>
        <Link href="keranjang">
          <a>Keranjang</a>
        </Link>
      </nav>
      <style jsx>
        {`
          nav {
            padding: 24px 32px;
          }
          nav a {
            background-color: rgba(200, 200, 200, 0.2);
            padding: 16px 32px;
            border-radius: 6px;
            transition: all 0.25s ease;
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
