import './App.css';
import React, { useState } from 'react';

function App() {
  const [items, setItems] = useState([]);
  const [item, setItem] = useState('');
  const [amount, setAmount] = useState('');
  const totalAmount = items.reduce((total, item) => total + parseInt(item.amount), 0);

  const handleAdd = (e) => {
    e.preventDefault();
    setItems([...items, { item, amount }]);
    setItem('');
    setAmount(0);
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>18th Feb - Sunday</h1>
      </header>
      <div className='container'>
        <div className='ledger'>
          <div className='items'>
            {items.map((item, index) => (
              <p className='item' key={index}>{item.item} - {item.amount}</p>
            ))}
          </div>
          <div className='input-container'>
            <p>Total: {totalAmount}</p>
            <form onSubmit={handleAdd} className='input-form'>
              <input 
                type='text' 
                placeholder='Enter item' 
                className='item-input'
                value={item}
                onChange={e => setItem(e.target.value)}
              />
              <input 
                type='number' 
                placeholder='Enter amount' 
                className='amount-input'
                value={amount}
                onChange={e => setAmount(e.target.value)}
              />
              <button type='submit'>Add</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
