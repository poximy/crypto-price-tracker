import { FC } from "react";
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
  return (
    <div className="App">
    </div>
  );
};

export default App;
