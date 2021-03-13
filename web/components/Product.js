import React from "react";

export default function Product(props) {
  return (
    <>
      <div className="product">
        <img src={props.image_url} alt="Product" />
        <div className="product-desc">
          <h3>{props.name}</h3>
          <p>Rp {props.price}</p>
        </div>
      </div>
      <style jsx>{`
        .product {
          width: 200px;
          border-radius: 6px;
          background-color: #fff;
          box-shadow: 2px 2px 10px 0 rgba(0, 0, 0, 0.15);
          cursor: pointer;
          overflow: hidden;
        }
        .product img {
          width: 100%;
          height: 200px;
          object-fit: cover;
        }
        .product-desc {
          padding: 8px 32px;
        }
      `}</style>
    </>
  );
}
