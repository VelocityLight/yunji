import * as React from "react";
import { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from "react-query";
import moment from 'moment';
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchTags, fetchServices } from "../api/filter_api"
import MultiSelector from "../components/MultiSelector"

import { Select } from 'antd';
import Selector from "../components/SingleSelector"
import { Divider, Tag, DatePicker } from 'antd';

import { fetchBillingByTagAndService } from "../api/billing_api"


const { Option } = Select;
const { RangePicker } = DatePicker;

export function getSelectorValues(tagResp) {
  return tagResp.map(tag => tag.name)
}

const BillingTrendPage = () => {
  const [selectedTags, setSelectedTags] = useState([])
  const [selectedService, setSelectedService] = useState([])
  const [inputData, setInputData] = useState([]);
  const onSelectTags = (values) => {
    setSelectedTags(values)
  }
  const onSelectService = (values) => {
    setSelectedService(values)
  }

  const tagQuery = useQuery(
    ["billingTrendTags", selectedTags],
    () => fetchTags(),
    {
      keepPreviousData: true,
      staleTime: 50,
    }

  )

  const serviceQuery = useQuery(
    ["billingTrendService", selectedService],
    () => fetchServices(),
    {
      keepPreviousData: true,
      staleTime: 50,
    }

  )

  useEffect(() => {
    fetchBillingByTagAndService({ tags: selectedTags, service: selectedService })
      .then((resp) => setInputData(resp.body == undefined ? [] : resp.body));
  }, [selectedTags, selectedService]);


  if (tagQuery.isLoading || serviceQuery.isLoading) {
    return <p>Loading...</p>;
  }
  if (tagQuery.error || serviceQuery.error) {
    return <p>Error: </p>;
  }

  const tagSelectors = getSelectorValues(tagQuery.data)
  const serviceSelectors = getSelectorValues(serviceQuery.data)
  // setInputData(query.data.body == undefined ? [] : query.data.body)
  // // const inputData = query.data.body == undefined ? [] : query.data.body
  const dateFormat = 'YYYYMMDD';
  return (
    <MyLayout>
      <div>
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
          color="cyan">Tags</Tag>
        <MultiSelector key="tags" items={tagSelectors} onSelect={onSelectTags} />
        <Tag style={{
          margin: '5px',
        }}
          color="cyan">Service</Tag>
        <MultiSelector key="services" items={serviceSelectors} onSelect={onSelectService} />


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
      <div>
        <PileBarChart data={inputData} />
      </div>
    </MyLayout >
  );
};

export default BillingTrendPage;
