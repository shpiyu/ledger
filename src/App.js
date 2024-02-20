import './App.css';
import React, { useState } from 'react';

function App() {
  const [items, setItems] = useState([]);
  const [item, setItem] = useState('');
  const [amount, setAmount] = useState('');
  const totalAmount = items.reduce((total, item) => total + parseFloat(item.amount), 0);

  const [date, setDate] = useState(new Date());

  const formatDate = (date) => {
    const day = date.getDate();
    const month = date.toLocaleDateString('default', { month: 'short' });
    const weekday = date.toLocaleDateString('default', { weekday: 'long' });
    const suffix = (day % 10 === 1 && day !== 11)
      ? 'st'
      : (day % 10 === 2 && day !== 12)
      ? 'nd'
      : (day % 10 === 3 && day !== 13)
      ? 'rd'
      : 'th';
    return `${weekday} - ${day}${suffix} ${month}`;
  }
  
  /**
   * Callback for add button. Adds item to list and resets input fields.
   * @param {*} e 
   */
  const handleAdd = (e) => {
    e.preventDefault();
    setItems([...items, { item, amount }]);
    setItem('');
    setAmount(0);
  }

  /**
   * Callback for delete button. Deletes item from list.
   * @param {*} e 
   */
  const handleDelete = (e) => {
    e.preventDefault();
    const item = e.target.value;
    for (let i = 0; i < items.length; i++) {
      if (items[i].item === item) {
        const newItems = items.slice();
        newItems.splice(i, 1);
        setItems(newItems);
        break;
      }
    }
  }

  /**
   * Callback for previous day button. Sets date to previous day.
   */
  const handlePrevDay = () => {
    setDate(prevDate => {
      const newDate = new Date(prevDate);
      newDate.setDate(prevDate.getDate() - 1);
      return newDate;
    });
  }

  /**
   * Callback for next day button. Sets date to next day.
   */
  const handleNextDay = () => {
    setDate(currDate => {
      const newDate = new Date(currDate);
      newDate.setDate(currDate.getDate() + 1);
      return newDate;
    });
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1 className='header'>
          <span className='nav-btn' onClick={handlePrevDay}>&lt;</span>
            {formatDate(date)}
          <span className='nav-btn' onClick={handleNextDay}>&gt;</span>
        </h1>
      </header>
      <div className='container'>
        <div className='ledger'>
          <div className='items'>
            {items.map((item, index) => (
              <div className='item' key={index}>
                <p>{item.item} - {item.amount}</p>
                <button value={item.item} onClick={handleDelete}>Delete</button>
              </div>
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
