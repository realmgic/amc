import React from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';

import VisibleViewClusterConnectionModal from 'containers/cluster/VisibleViewClusterConnectionModal';
import ClusterOverview from 'components/cluster/ClusterOverview';
import EditClusterConnectionModal from 'components/cluster/EditClusterConnectionModal';
import { CLUSTER_ACTIONS } from 'classes/entityActions';
import Tabs from 'components/Tabs';
import ClusterLatency from 'components/cluster/ClusterLatency';
import ClusterNodesConfig from 'components/cluster/ClusterNodesConfig';

// ClusterDashboard handles all the views for the cluster.
// It is also responsible for changing between different views
// of the cluster.
//
// The state of the view, view type is maintained outside of the component.
// Hence there are callbacks to the parent component to change these view states.
class ClusterDashboard extends React.Component {
  constructor(props) {
    super(props);

    this.views = [CLUSTER_ACTIONS.Overview, CLUSTER_ACTIONS.Latency, CLUSTER_ACTIONS.Configuration];

    this.onViewSelect = this.onViewSelect.bind(this);
  }

  onViewSelect(view) {
    const { clusterID } = this.props;
    this.props.onViewSelect(clusterID, view);
  }

  render() {
    const { clusterID, view, onSelectNode }  = this.props;
    const { name, seeds } = this.props.cluster;

    let dashboard;
    if (view === CLUSTER_ACTIONS.Overview ) {
      dashboard = <ClusterOverview clusterID={clusterID} onSelectNode={onSelectNode} />;

    } else if (view === CLUSTER_ACTIONS.Latency) {
      dashboard = <ClusterLatency clusterID={clusterID} />

    } else if (view === CLUSTER_ACTIONS.Configuration) {
      dashboard = <ClusterNodesConfig clusterID={clusterID} />
    }

    return (
      <div>
        <Tabs names={this.views} selected={view} onSelect={this.onViewSelect}/>

        {dashboard}
      </div>
    );
  }
}

ClusterDashboard.PropTypes = {
  clusterID: PropTypes.string.isRequired,
  // the selected cluster
  cluster: PropTypes.object, 
  // the view of the cluster
  view: PropTypes.string,

  // callback to select a node
  // onSelectNode(clusterID, nodeHost)
  onSelectNode: PropTypes.func,
};

export default ClusterDashboard;

