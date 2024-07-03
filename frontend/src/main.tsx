// import './../wdyr'; // <--- first import
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)



// import 'react-hot-loader';
// import {hot} from 'react-hot-loader/root';

// import React from 'react';
// import ReactDOM from 'react-dom';
// // ...
// import {App} from './app';
// // ...
// const HotApp = hot(App);
// // ...
// ReactDOM.render(<HotApp/>, document.getElementById('root'));