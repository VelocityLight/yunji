import * as React from "react";
import { useEffect, useState } from 'react';
import moment from 'moment';
import MyLayout from "../layout/Layout";
import MultiSelector from "../components/MultiSelector"

import Selector from "../components/SingleSelector"
import { Select } from 'antd';
import { Divider, Tag, DatePicker } from 'antd';

import { useQuery, useQueryClient } from "react-query";
import PieChart from "../graphs/PieChart";
import { fetchTeamCosts } from "../api/team_api";


const { Option } = Select;
const { RangePicker } = DatePicker;

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


  const dateFormat = 'YYYYMMDD';
  return (
    <MyLayout>
      <div       >
        <Tag style={{
          margin: '5px',
        }}
          color="cyan">Time</Tag>
        <RangePicker
          defaultValue={[moment('20220801', dateFormat), moment('20220831', dateFormat)]}
          disabled={[false, false]}
        />

        <Tag style={{
          margin: '5px',
        }}
          color="cyan">Cloud</Tag>
        <Select
          defaultValue="AWS"
          style={{
            width: "10%",
          }}
        >
          <Option key="AWS">AWS</Option>
          <Option key="GCP" disabled>GCP</Option>
          <Option key="AZURE" disabled>AZURE</Option>
          <Option key="Ali Cloud" disabled>Ali Cloud</Option>
        </Select>
      </div>

      <PieChart data={pieData} type_key="team" value_key={"cost"} />
    </MyLayout>
  );
};

export default TeamPage;
