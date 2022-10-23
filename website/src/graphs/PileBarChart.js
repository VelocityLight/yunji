import React from 'react';
import { Column } from '@ant-design/plots';
import { each, groupBy } from '@antv/util';
import { useEffect, useState } from 'react';
import { useQuery } from "react-query";
import { fetchBillingByTagAndService } from "../api/billing_api"

const PileBarChart = ({ tags, service }) => {
  const [data, setData] = useState([]);

  useEffect(() => {
    fetchBillingByTagAndService({ tags: tags, service: service }).then((resp) => setData(resp.body));
  }, [tags, service]);

  // const annotations = [];

  // const query = useQuery(
  //   ["billingTrend", tags, service],
  //   () =>
  //     fetchBillingByTagAndService({ tags: tags, service: service }).then,
  //   {
  //     keepPreviousData: true,
  //     staleTime: 20000,
  //   }
  // );

  // each(groupBy(query.body, 'time'), (data, k) => {
  //   const value = data.reduce((a, b) => a + b.value, 0);
  //   annotations.push({
  //     type: 'text',
  //     position: [k, value],
  //     content: `${value}`,
  //     style: {
  //       textAlign: 'center',
  //       fontSize: 14,
  //       fill: 'rgba(0,0,0,0.85)',
  //     },
  //     offsetY: -10,
  //   });
  // });

  const config = {
    data,
    isStack: true,
    xField: 'time',
    yField: 'cost',
    seriesField: 'service',
    label: {
      // 可手动配置 label 数据标签位置
      position: 'middle', // 'top', 'bottom', 'middle'
    },
    interactions: [
      {
        type: 'active-region',
        enable: false,
      },
    ],
    connectedArea: {
      style: (oldStyle, element) => {
        return {
          fill: 'rgba(0,0,0,0.25)',
          stroke: oldStyle.fill,
          lineWidth: 0.5,
        };
      },
    },
    // annotations,
  };

  // console.log(query.body)

  // if (query.isLoading) {
  //   return <p>Loading...</p>;
  // }
  // if (query.error) {
  //   return <p>Error: {query.error.message}</p>;
  // }

  return <Column {...config} />;
};

export default PileBarChart;
