import React from 'react';
import { render } from 'react-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

import VisibleClusterDashboard from 'containers/cluster/VisibleClusterDashboard';
import VisibleUDFView from 'containers/udf/VisibleUDFView';
import VisibleUDFOverviewDashboard from 'containers/udf/VisibleUDFOverviewDashboard'
import VisibleSetsOverview from 'containers/set/VisibleSetsOverview';
import VisibleSetView from 'containers/set/VisibleSetView';
import VisibleIndexView from 'containers/index/VisibleIndexView';
import VisibleIndexesOverview from 'containers/index/VisibleIndexesOverview';

import NodeDashboard from 'components/node/NodeDashboard';
import NodesOverview from 'components/node/NodesOverview';
import NamespaceDashboard from 'components/namespace/NamespaceDashboard';
import Welcome from 'components/Welcome';

import { VIEW_TYPE } from 'classes/constants';
import { CLUSTER_ACTIONS } from 'classes/entityActions';

class MainDashboard extends React.Component {
  constructor(props) {
    super(props);

    this.onChangeView = this.onChangeView.bind(this);
  }

  onChangeView(toView) {
    const { selectedEntityPath, view } = this.props.currentView;
    if (toView === view)
      return;

    this.props.onSelectPath(selectedEntityPath, toView);
  }

  render() {
    const { currentView, isClusterConnected, clusterName } = this.props;
    const { entities, viewType, view } = currentView;
    const { clusterID, nodeHost, namespaceName, setName, udfName } = currentView;

    let dashboard;

    if (clusterID && !isClusterConnected) {
      if (viewType === VIEW_TYPE.CLUSTER && view === CLUSTER_ACTIONS.Edit)
        dashboard = <VisibleClusterDashboard />
      else
        dashboard = <h4 style={{marginTop: 20}}> Please connect to {`"${clusterName}"`} to continue </h4>;

    } else if (viewType === VIEW_TYPE.NODE_OVERVIEW) {
      dashboard = <NodesOverview clusterID={clusterID} />

    } else if (viewType === VIEW_TYPE.NODE) {
      dashboard = <NodeDashboard clusterID={clusterID} nodeHost={nodeHost} view={view} onViewSelect={this.onChangeView}/>

    } else if (viewType === VIEW_TYPE.NAMESPACE) {
      dashboard = <NamespaceDashboard clusterID={clusterID} nodeHost={nodeHost} namespaceName={namespaceName}
                    view={view} onViewSelect={this.onChangeView}/>

    } else if (viewType === VIEW_TYPE.SET) {
      dashboard = <VisibleSetView />

    } else if (viewType === VIEW_TYPE.SET_OVERVIEW) {
      dashboard = <VisibleSetsOverview />

    } else if (viewType === VIEW_TYPE.INDEXES_OVERVIEW) {
      dashboard = <VisibleIndexesOverview />
        
    } else if (viewType === VIEW_TYPE.INDEX) {
      dashboard = <VisibleIndexView />
        
    } else if (viewType === VIEW_TYPE.UDF) {
      dashboard = <VisibleUDFView />

    } else if (viewType === VIEW_TYPE.UDF_OVERVIEW) {
      dashboard = <VisibleUDFOverviewDashboard />

    } else if (viewType === VIEW_TYPE.CLUSTER) {
      dashboard = <VisibleClusterDashboard />

    } else if (viewType === VIEW_TYPE.START_VIEW) {
      dashboard = <Welcome />;

    } else {
      dashboard = <div className="as-centerpane-header"> {(view ? view : '') + ' ' + viewType} </div>;

    }

    return (
      <div>
        {dashboard}
      </div>
      );
  }
}

MainDashboard.PropTypes = {
  // current view state
  currentView: PropTypes.object,
  // select a path
  // onSelectPath(path, view)
  onSelectPath: PropTypes.func,
  // if a cluster entity is selected 
  // is the cluster connected
  isClusterConnected: PropTypes.bool,
  // name of the cluster the dashboard is in
  clusterName: PropTypes.bool,
};

export default MainDashboard;
