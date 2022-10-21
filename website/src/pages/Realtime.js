import * as React from "react";
import { useEffect, useState } from 'react';
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";

const RealtimePage = () => {

  const [data, setData] = useState([]);

  useEffect(() => {
    asyncFetch();
  }, []);

  const asyncFetch = () => {
    fetch('https://gw.alipayobjects.com/os/antfincdn/8elHX%26irfq/stack-column-data.json')
      .then((response) => response.json())
      .then((json) => setData(json))
      .catch((error) => {
        console.log('fetch data failed', error);
      });
  };

  console.log(data)

  return (
    <MyLayout>
      <PileBarChart data={data} xfiled_key={"year"} yfiled_key={"value"} serie_key="type" />
    </MyLayout>
  );
};

export default RealtimePage;
