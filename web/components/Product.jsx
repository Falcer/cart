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
          margin: 24px;
          width: 200px;
          border-radius: 6px;
          background-color: #fff;
          box-shadow: 5px 5px 10px 0 rgba(0, 0, 0, 0.05);
          cursor: pointer;
          transform: scale(1);
          transition: all 0.25s ease;
          overflow: hidden;
        }
        .product:hover {
          transform: scale(1.1);
          box-shadow: 5px 5px 30px 0 rgba(0, 0, 0, 0.35);
        }
        .product:hover img {
          transform: scale(1.0075);
        }
        .product img {
          width: 100%;
          height: 200px;
          object-fit: cover;
          transform: scale(1);
          overflow: hidden;
          transition: all 0.25s ease;
        }
        .product-desc {
          padding: 8px 32px;
        }
      `}</style>
    </>
  );
}
