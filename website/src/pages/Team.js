import * as React from "react";
import { useEffect, useState } from 'react';
import MyLayout from "../layout/Layout";

import { useQuery, useQueryClient } from "react-query";
import PieChart from "../graphs/PieChart";
import { fetchTeamCosts } from "../api/team_api";

const TeamPage = () => {

  const query = useQuery(
    ["team"],
    () =>
      fetchTeamCosts({}),
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

  const pieData = query.data

  return (
    <MyLayout>
      <PieChart data={pieData} type_key="team" value_key={"cost"} />
    </MyLayout>
  );
};

export default TeamPage;
