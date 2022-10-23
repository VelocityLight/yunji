import * as React from "react";
import { useEffect, useState } from 'react';
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchRealtime } from "../api/billing_api"

const RealtimePage = () => {
  const [inputData, setInputData] = useState([]);

  useEffect(() => {
    const interval = setInterval(() => {
      console.log('This will run every 10 second!');
      fetchRealtime({})
        .then((resp) => {
          var tmp = resp.body == undefined ? [] : resp.body
          setInputData(tmp)
        });

    }, 10000);
    return () => clearInterval(interval);
  }, []);

  console.log(inputData)

  return (
    <MyLayout>
      <PileBarChart data={inputData} xfiled={"time"} yfiled={"cnt"} serie="service" />
    </MyLayout>
  );
};

export default RealtimePage;
