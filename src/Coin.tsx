import { FC } from "react";

interface CoinStats {
  image: string;
  name: string;
  symbol: string;
  price: number;
  volume: number;
  change: number;
  marketCap: number;
}

const CoinData: FC<CoinStats> = ({
  image,
  name,
  symbol,
  price,
  volume,
  change,
  marketCap,
}) => {
  return (
    <div className="coin-container">
      <div className="coin-row">
        <div className="coin">
          <img src={image} alt="cryto" />
          <h1>{name}</h1>
          <p className="coin-symbol">{symbol}</p>
        </div>
        <div className="coin-data">
          <p className="coin-price">${price}</p>
          <p className="coin-volume">${volume.toLocaleString()}</p>
          <p className={`coin-percent ${change < 0 ? "red" : "green"}`}>
            {change.toFixed(2)}%
          </p>
          <p className="coin-marketcap">
            Mkt Cap: ${marketCap.toLocaleString()}
          </p>
        </div>
      </div>
    </div>
  );
};

export default CoinData;
