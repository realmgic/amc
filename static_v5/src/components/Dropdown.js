import React from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types'
import { nextNumber } from '../classes/Util';
import { Dropdown as ReactDropdown, DropdownMenu, DropdownItem } from 'reactstrap';

import 'bootstrap/dist/css/bootstrap.css';

// A custom dropdown component
// see https://reactstrap.github.io/components/dropdowns/
class Dropdown extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      dropdownOpen: false
    };

    this.toggle = this.toggle.bind(this);
    this.onOptionClick = this.onOptionClick.bind(this);
    this.onLabelClick = this.onLabelClick.bind(this);
  }

  toggle() {
    this.setState({
      dropdownOpen: !this.state.dropdownOpen
    });
  }

  onOptionClick(option) {
    this.toggle();
    this.props.onOptionSelect(option);
  }

  onLabelClick() {
    this.props.onLabelClick();
  }

  render() {
    const toggleStyle = {
      marginLeft: 5
    };
    return (
      <ReactDropdown isOpen={this.state.dropdownOpen} toggle={this.toggle} style={this.props.style}>
        <span onClick={this.onLabelClick}> {this.props.label} </span>
        <button onClick={this.toggle} style={toggleStyle} className="btn btn-sm dropdown-toggle"> </button>
        <DropdownMenu>
          {this.props.options.map((option) => {
             return <DropdownItem key={nextNumber()} onClick={(evt) => this.onOptionClick(option)}>
                      {typeof option === 'string' ? option : option.label}
                    </DropdownItem>
           })}
        </DropdownMenu>
      </ReactDropdown>
      );
  }
}

Dropdown.PropTypes = {
  // label of the dropdown
  label: PropTypes.string.isRequired,
  // callback when an option is selected
  // onOptionSelect(option)
  onOptionSelect: PropTypes.func,
  // callback when the label is clicked
  // onLabelClick(void)
  onLabelClick: PropTypes.func,
  // style applied to the entire dropdown
  style: PropTypes.object,

  // the options to display
  // [string | object { label: 'option_label'}]
  options: PropTypes.arrayOf(function(propValue, key, componentName, location, propFullname) {
    const prop = propValue[key];
    const type = typeof prop;
    if (type === 'string') return;
    if (type === 'object' && typeof prop.label === 'string') return;
    return new Error(
      'Invalid prop `' + propFullName + '` supplied to' +
      ' `' + componentName + '`. Validation failed.'
    );
  }).isRequired,
};

export default Dropdown;