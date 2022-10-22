import { Select } from 'antd';
import React from 'react';
const { Option } = Select;

const Selector = ({ items = [], onSelect = (newValue) => { return newValue } }) => {
  const options = [];
  for (let i = 0; i < items.length; i++) {
    options.push(<Option key={i + items[i]}>{items[i]}</Option>);
  }

  return (
    <Select
      defaultValue=""
      style={{
        width: "25%",
      }}
      onChange={onSelect}
    >
      <Option key="">-</Option>
      {options}
    </Select>
  );
};

export default Selector;

