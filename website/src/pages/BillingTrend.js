import * as React from "react";
import { useEffect, useState } from 'react';
import { useQuery, useQueryClient } from "react-query";
import MyLayout from "../layout/Layout";
import PileBarChart from "../graphs/PileBarChart";
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
  const onSelectService = (values) => {
    setSelectedService(values)
  }

  const tagQuery = useQuery(
    ["billingTrend", selectedTags],
    () => fetchTags(),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  )

  const serviceQuery = useQuery(
    ["billingTrend", selectedService],
    () => fetchServices(),
    {
      keepPreviousData: true,
      staleTime: 5000,
    }
  )

  if (tagQuery.isLoading || serviceQuery.isLoading) {
    return <p>Loading...</p>;
  }
  if (tagQuery.error || serviceQuery.error) {
    return <p>Error: {tagQuery.error.message}</p>;
  }

  const tagSelectors = getSelectorValues(tagQuery.data)
  const serviceSelectors = getSelectorValues(serviceQuery.data)

  return (
    <MyLayout>
      <div>
        <Tag color="cyan">Tags</Tag>
        <MultiSelector key="tags" items={tagSelectors} onSelect={onSelectTags} />
        <Tag color="cyan">Service</Tag>
        <MultiSelector key="services" items={serviceSelectors} onSelect={onSelectService} />
      </div>
      <div>
        <PileBarChart tags={selectedTags} service={selectedService} />
      </div>
    </MyLayout>
  );
};

export default BillingTrendPage;
