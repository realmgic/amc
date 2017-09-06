import React from 'react';
import { render } from 'react-dom';
import PropTypes from 'prop-types';

import { isPermissibleAction, INDEX_ACTIONS } from 'classes/entityActions';
import { getIndex, deleteIndex } from 'api/index';
import { VIEW_TYPE } from 'classes/constants';
import IndexesTable from 'components/index/IndexesTable';
import Spinner from 'components/Spinner';
import AlertModal from 'components/AlertModal';
import { whenClusterHasCredentials } from 'classes/security';

import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from 'reactstrap';

// IndexView diplays a view of the index
class IndexView extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      data: null,

      canDelete: false,

      deleteShowConfirm: false,
      deleteInProgress: false,
      deleteSuccessfull: null,
      deleteErrorMsg: '',
    };

    this.onShowConfirm = this.onShowConfirm.bind(this);
    this.onDeleteSuccess = this.onDeleteSuccess.bind(this);
    this.onDeleteIndex = this.onDeleteIndex.bind(this);
  }

  componentDidMount() {
    const { clusterID } = this.props;
    this.setPermissions(clusterID);
  }

  setPermissions(clusterID) {
    whenClusterHasCredentials(clusterID, () => {
      const canDelete = isPermissibleAction(INDEX_ACTIONS.Delete, clusterID, VIEW_TYPE.INDEX);
      this.setState({
        canDelete: canDelete
      });
    });
  }

  onShowConfirm() {
    this.setState({
      deleteShowConfirm: true,
    });
  }

  onDeleteSuccess() {
    const { clusterID, indexName } = this.props;
    this.props.onDeleteSuccess(clusterID, indexName);
  }

  onDeleteIndex() {
    const { clusterID, indexName } = this.props;
    const { namespace, set } = this.state.data;

    this.setState({
      deleteInProgress: true
    });
    deleteIndex(clusterID, namespace, set, indexName)
      .then(() => {
        this.setState({
          deleteInProgress: false,
          deleteSuccessfull: true
        });
        window.setTimeout(() => this.onDeleteSuccess(), 2000);
      })
      .catch((msg) => {
        this.setState({
          deleteInProgress: false,
          deleteSuccessfull: false,
          deleteErrorMsg: msg || 'Failed to delete index'
        });
      });
  }

  fetchData(clusterID, indexName) {
    this.setState({
      data: null,
    });

    getIndex(clusterID, indexName)
      .then((index) => {
        this.setState({
          data: index
        });
      })
      .catch((message) => {
        console.error(message);
      });
  }

  componentWillMount() {
    const {clusterID, indexName} = this.props;
    this.fetchData(clusterID, indexName);
  }

  componentWillReceiveProps(nextProps) {
    let isSame = true;
    ['clusterID', 'indexName'].forEach((p) => {
      if (nextProps[p] !== this.props[p])
        isSame = false;
    });

    if (isSame)
      return;

    const {clusterID, indexName} = nextProps;
    this.setPermissions(clusterID);
    this.fetchData(clusterID, indexName);
  }

  renderDeleteModal() {
    const { deleteShowConfirm, deleteInProgress, deleteSuccessfull, deleteErrorMsg } = this.state;

    const onCancelModal = () => {
      this.setState({
        deleteShowConfirm: false
      });
    };

    if (!deleteShowConfirm)
      return null;

    const disabled = deleteInProgress || deleteSuccessfull;
    if (!deleteInProgress && deleteSuccessfull === true) {
      const message = `Successfully deleted ${this.props.indexName}`;
      return (
        <AlertModal header="Success" message={message} type="success" />
      );
    }

    return (
      <Modal isOpen={true} toggle={() => {}}>
        <ModalHeader className="alert-danger"> Confirm </ModalHeader>
        <ModalBody>  Delete {this.props.indexName} ?  </ModalBody>
        <ModalFooter>
          {!deleteInProgress && deleteSuccessfull === false &&
            deleteErrorMsg}
          {deleteInProgress &&
           <span> <Spinner /> Deleting ... </span>}
          <Button disabled={disabled} color="danger" onClick={this.onDeleteIndex}>Confirm</Button>
          <Button disabled={disabled} color="secondary" onClick={onCancelModal}>Cancel</Button>
        </ModalFooter>
      </Modal>
    );
  }

  render() {
    const {view} = this.props;

    const indexes = [];
    if (this.state.data)
      indexes.push(this.state.data);

    const { deleteInProgress, canDelete } = this.state;
    const {clusterID, indexName} = this.props;

    return (
      <div>
        {this.renderDeleteModal()}

        <div className="row">
          <div className="col-xl-12 as-section-header">
            {`Index - ${indexName}`} 

            {canDelete &&
            <Button className="float-right" disabled={deleteInProgress} color="danger" size="sm" onClick={this.onShowConfirm}> 
              <i className="fa fa-trash"></i>
              Delete 
            </Button>}
          </div>
        </div>

        <IndexesTable indexes={indexes} header={indexName} initiallyExpandAll={true}/>

      </div>
    );
  }
}

IndexView.PropTypes = {
  clusterID: PropTypes.string.isRequired,
  indexName: PropTypes.string.isRequired,

  // callback on successfull delete
  // onDeleteSuccess(clusterID, indexName)
  onDeleteSuccess: PropTypes.func,
};


export default IndexView;

