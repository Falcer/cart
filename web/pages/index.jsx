import React from "react";
import axios from "axios";
import Navbar from "../components/Navbar";
import Product from "../components/Product";

export default function Home() {
  const [products, setProducts] = React.useState([]);
  const [loading, setLoading] = React.useState(false);

  React.useEffect(() => {
    setLoading(true);
    axios
      .get("http://54.169.75.182:8080/api/v1/products")
      .then((result) => {
        setProducts(result.data.data);
      })
      .catch((err) => {
        console.log(`Got some error ${err}`);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  return (
    <>
      <Navbar />
      {loading ? <h1>Loading . . .</h1> : null}
      <section>
        {products.length === 0 && !loading ? (
          <h1>Products Empty</h1>
        ) : (
          products.map((item, key) => (
            <Product
              key={key}
              name={item.name}
              image_url={item.image_url}
              price={item.price}
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
