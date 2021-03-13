import React from "react";
import Navbar from "../components/Navbar";
import Product from "../components/Product";

export default function Home() {
  const [products, setProducts] = React.useState([]);
  return (
    <>
      <Navbar />
      <section>
        {products.length === 0 ? (
          <h1>Products Empty</h1>
        ) : (
          products.map((item, key) => (
            <Product
              key={key}
              name="ROG"
              image_url="https://images.unsplash.com/photo-1615148536759-8fae933d8ff1?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=1080&fit=max"
              price={2_000_000}
            />
          ))
        )}
      </section>
      <style jsx>{`
        section {
          margin: 56px 0 0;
          padding: 0 96px;
          display: flex;
          flex-wrap: wrap;
          justify-content: center;
        }
      `}</style>
    </>
  );
}
