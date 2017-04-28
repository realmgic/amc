import React from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import VisibleEntityTree from '../containers/VisibleEntityTree';
import VisibleClusterToolbar from '../containers/VisibleClusterToolbar';
import VisibleAddClusterModal from '../containers/VisibleAddClusterModal';

class ClusterConnections extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    let content;
    if (this.props.isFetching) {
      content = <h5> Loading ... </h5>;
    } else {
      content = (
        <div>
          <VisibleClusterToolbar />
          <VisibleEntityTree />
          {this.props.displayAddCluster &&
           <VisibleAddClusterModal />}
        </div>
      );
    }
    return content;
  }
}

ClusterConnections.PropTypes = {
  // true when in the process of adding a new connection
  displayAddCluster: PropTypes.bool,
  isFetching: PropTypes.bool,
};

export default ClusterConnections;

