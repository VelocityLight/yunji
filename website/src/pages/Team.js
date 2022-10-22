import * as React from "react";
import { useEffect, useState } from 'react';
import MyLayout from "../layout/Layout";
import PieChart from "../graphs/PieChart";

const TeamPage = () => {

  const [data, setData] = useState([
    {
      type: '分类一',
      value: 100,
    },
    {
      type: '分类二',
      value: 200,
    },
    {
      type: '分类三',
      value: 300,
    },
    {
      type: '分类四',
      value: 100,
    },
    {
      type: '分类五',
      value: 100,
    },
    {
      type: '其他',
      value: 200,
    },
  ]);

  return (
    <MyLayout>
      <PieChart data={data} type_key="type" value_key={"value"} />
    </MyLayout>
  );
};

export default TeamPage;
