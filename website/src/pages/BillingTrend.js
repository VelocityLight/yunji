import * as React from "react";
import { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from "react-query";
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
import { fetchBillingByTagAndService } from "../api/billing_api"
import { fetchTags, fetchServices } from "../api/filter_api"
import MultiSelector from "../components/MultiSelector"
import Selector from "../components/SingleSelector"
import { Divider, Tag } from 'antd';

export function getSelectorValues(tagResp) {
  return tagResp.map(tag => tag.name)
}

const BillingTrendPage = () => {
  const [selectedTags, setSelectedTags] = useState([])
  const [selectedService, setSelectedService] = useState([])
  const onSelectTags = (values) => {
    setSelectedTags(values)
  }
  const onSelectService = (key, value) => {
    setSelectedService(value.value)
  }

  const tagQuery = useQuery(
    ["billingTrend", selectedService],
    () => fetchTags({}),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  )

  const serviceQuery = useQuery(
    ["billingTrend", selectedTags],
    () => fetchServices({ selectedTags }),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  )

  const query = useQuery(
    ["billingTrend", selectedTags, selectedService],
    () =>
      fetchBillingByTagAndService({ tags: selectedTags, service: selectedService }),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  );

  if (query.isLoading || tagQuery.isLoading || serviceQuery.isLoading) {
    return <p>Loading...</p>;
  }
  if (query.error || tagQuery.error || serviceQuery.error) {
    return <p>Error: {query.error.message}</p>;
  }

  const inputData = query.data.body
  const tagSelectors = getSelectorValues(tagQuery.data)
  const serviceSelectors = getSelectorValues(serviceQuery.data)

  return (
    <MyLayout>
      <div>
        <Tag color="cyan">Tags</Tag>
        <MultiSelector items={tagSelectors} onSelect={onSelectTags} />
        <Tag color="cyan">Service</Tag>
        <Selector items={serviceSelectors} onSelect={onSelectService} />
      </div>
      <div>
        <PileBarChart data={inputData} xfiled_key={"time"} yfiled_key={"cost"} serie_key="tag" />
      </div>
    </MyLayout>
  );
};

export default BillingTrendPage;
