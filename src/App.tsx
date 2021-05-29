import { FC, useEffect, useState, ChangeEvent } from "react";
import axios from "axios";

import Coin from "./apiModel";

import "./App.css";

const coinData = async () => {
  try {
    const res = await axios.get(
      "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false"
    );
    const data: Coin[] = res.data;
    return res.data ? data : "ERROR";
  } catch (err) {
    return "ERROR";
  }
};

const App: FC = () => {
  const [coins, setCoins] = useState<Coin[]>([]);
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    (async () => {
      const data = await coinData();
      data !== "ERROR" ? setCoins(data) : console.log(data);
    })();
  }, []);

  const searchChange = (e: ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
  };

  const filteredCoins = coins.filter((coin) => {
    coin.name.toLowerCase().includes(search.toLowerCase());
  });

  return (
    <div className="coin-app">
      <div className="coin-search">
        <h1 className="coin-text">Search a currency</h1>
        <form>
          <input
            className="coin-input"
            type="text"
            placeholder="search"
            onChange={searchChange}
          />
        </form>
      </div>
    </div>
  );
};

export default App;
