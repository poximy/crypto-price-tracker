const socket = new WebSocket('ws://' + location.host + '/ws');
socket.onopen = () => console.log('WebSocket opened!');
socket.onerror = (err) => {
  alert('Error with the web socket conection! See the console');
  console.error(err);
};

interface CoinGeko {
	name: string,
	image: string,
	symbol: string,
	current_price: number,
	price_change_24h: number,
	price_change_percentage_24h: number
}

function renderCard(data: CoinGeko) {
  return `
      <div
        class="flex h-48 flex-col items-center justify-center gap-2 border
        border-neutral-200 bg-neutral-100 p-2"
        data-name=${data.name}"
      >
        <img
          src=${data.image}
          alt="crypto"
          class="mx-auto aspect-square w-12 sm:w-16"
          loading="lazy"
        />
        <div class="flex w-64 justify-evenly gap-2 text-lg sm:w-48">
          <p class="truncate">${data.name}</p>
          <p class="font-bold uppercase">${data.symbol}</p>
        </div>
        <div class="flex w-64 justify-evenly font-mono sm:w-48">
          <p>${data.current_price}$</p>
          <p>${data.price_change_percentage_24h}%</p>
        </div>
      </div>
  `;
}

window.addEventListener('DOMContentLoaded', () => {
	// @ts-ignore
  const container: HTMLDivElement = document.getElementById('container')

  socket.onmessage = (ev) => {
    const data: CoinGeko[] = JSON.parse(ev.data);
    container.innerHTML = data.map(renderCard).join('\n');
  };
});
