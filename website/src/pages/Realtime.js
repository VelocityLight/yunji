import * as React from "react";
import { useEffect, useState } from 'react';
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchRealtime } from "../api/realtime_api"

const RealtimePage = () => {
  const [inputData, setInputData] = useState([]);

  useEffect(() => {
    const interval = setInterval(() => {
      console.log('This will run every 5 second!');
      fetchRealtime({})
        .then((resp) => {
          var tmp = resp.body == undefined ? [] : resp.body
          setInputData(tmp)
        });

    }, 5000);
    return () => clearInterval(interval);
  }, []);


  if (inputData.length == 0) {
    return <p>Loading...</p>;
  }

  console.log(inputData)

  return (
    <MyLayout>
      <PileBarChart data={inputData} xfiled="time" yfield="cnt" serie="service" />
    </MyLayout>
  );
};

export default RealtimePage;
