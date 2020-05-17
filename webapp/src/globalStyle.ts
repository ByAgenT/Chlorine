import { createGlobalStyle } from 'styled-components';

const AppStyle = createGlobalStyle`
  @font-face {
    font-family: 'Nunito';
    src: local('Nunito'), local('Nunito'),
    url('/fonts/Nunito-Light.ttf') format('truetype');
    font-weight: 300;
    font-style: normal;  
  }

  body {
    margin: 0;
    padding: 0;
    font-size: 10px;
    background-color: #222326;
    font-family: 'Nunito', 'Josefin Sans', -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",
      "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
      sans-serif;
  }
`;

export default AppStyle;
