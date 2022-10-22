import * as React from "react";
import { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from "react-query";
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchBillingByTagAndService } from "../api/billing_api"

const BillingTrendPage = () => {

  // const [inputData, setInputData] = useState([]);

  const query = useQuery(
    ["billingTrend", "tags", "service"],
    () =>
      fetchBillingByTagAndService({}),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  );

  if (query.isLoading) {
    return <p>Loading...</p>;
  }
  if (query.error) {
    return <p>Error: {query.error.message}</p>;
  }

  console.log(query.data)
  const inputData = query.data.body

  return (
    <MyLayout>
      <PileBarChart data={inputData} xfiled_key={"time"} yfiled_key={"cost"} serie_key="tag" />
    </MyLayout>
  );
};

export default BillingTrendPage;
