import * as React from "react";
import { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from "react-query";
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchBillingByTagAndService } from "../api/billing_api"

const BillingTrendPage = () => {

  const [data, setData] = useState([]);


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

  console.log(data)
  setData(query.data.body)

  return (
    <MyLayout>
      <PileBarChart data={data} xfiled_key={"time"} yfiled_key={"cost"} serie_key="tag" />
    </MyLayout>
  );
};

export default BillingTrendPage;
