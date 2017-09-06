import { getRoles as getRolesAPI } from 'api/clusterRoles';
import { getUsers as getUsersAPI } from 'api/clusterUsers';
import { getConnectionDetails, getLoggedInUser } from 'api/clusterConnections';
import { VIEW_TYPE } from 'classes/constants';

let IsInitialized = false;
// map of cluster id to the cluster properties
let ClusterProps = {};
// map of cluster id to the logged in user roles
let UserRoles = {};
// map of cluster id to the roles defined in the cluster
let ClusterRoles = {};
let IsAMCEnterprise = false;
// roles of the user in AMC
let AmcRoles = [];

function hasCredentials(clusterID) {
  const id = clusterID;
  return id in ClusterProps && 
         id in ClusterRoles &&
         id in UserRoles;
}

export function whenClusterHasCredentials(clusterID, fn) {
  const check = () => {
    if (hasCredentials(clusterID))
      fn();
    else
      window.setTimeout(check, 200);
  };
  check();
}

// secureCluster adds security functionality to the cluster
export function secureCluster(clusterID ) {
  // waiting a few seconds for the server to fetch details about the cluster
  window.setTimeout(() => {
    // fetch all roles of cluster
    getRolesAPI(clusterID)
      .then((roles) => {
        if (!IsInitialized)
          return;

        ClusterRoles[clusterID] = roles;
      });

    // get roles of logged in user
    getLoggedInUser(clusterID)
      .then((user) => {
        if (!IsInitialized)
          return;

        UserRoles[clusterID] = user.roles || [];
      });

    // get properties of cluster
    getConnectionDetails(clusterID)
      .then((cluster) => {
        if (!IsInitialized)
          return;

        ClusterProps[clusterID] = {
          isAuthenticated: true,
          isSecure: cluster.isSecurityEnabled,
        };
      });
  }, 5000);
}

export function removeCluster(clusterID) {
  const id = clusterID;
  delete UserRoles[id];
  delete ClusterProps[id];
  delete ClusterRoles[id];
}

export function removeAllClusters() {
  IsInitialized = false;

  let keys = [];
  keys = keys.concat(Object.keys(UserRoles));
  keys = keys.concat(Object.keys(ClusterProps));
  keys = keys.concat(Object.keys(ClusterRoles));

  keys.forEach((id) => removeCluster(id));
}

// init initializes the security for this session of the user
export function init(roles, isEnterprise) {
  IsInitialized = true;
  IsAMCEnterprise = isEnterprise;
  AmcRoles = roles;
}

function hasBasicAccess(clusterID, clusterProps = {}, isEnterprise = false, 
            amcRoles = []) {
  if (isEnterprise && !IsAMCEnterprise)
    return false;

  const props = ClusterProps[clusterID] || {};
  if (!isPropsSubset(clusterProps, props))
    return false;

  if (!isArrSubset(amcRoles, AmcRoles))
    return false;

  return true;
}
// returns true iff the current user satisfies the privileges for the cluster
// as specified in the arguments.
//
// dbRoles - the roles the user should have in the cluster
// clusterProps - the properties that the cluster should have
// isEnterprise - should this version of amc be enterprise
// amcRoles - the roles the amc user should have
export function canAccess(clusterID, dbRoles = [], clusterProps = {},
                  isEnterprise = false, amcRoles = []) {
  if (!hasBasicAccess(clusterID, clusterProps, isEnterprise, amcRoles))
      return false;

  const props = ClusterProps[clusterID] || {};
  if (props.isSecure) {
    const roles = getRolesForUser(clusterID);
    return hasRoles(roles, dbRoles);
  }

  return true;
}

// returns true iff the current user satisfies the privileges for the namespace
// as specified in the arguments.
export function canAccessNamespace(clusterID, namespace, dbRoles = [], 
                    clusterProps = {}, isEnterprise = false, amcRoles = []) {
  if (!hasBasicAccess(clusterID, clusterProps, isEnterprise, amcRoles))
      return false;

  const props = ClusterProps[clusterID] || {};
  if (props.isSecure) {
    const roles = getRolesForUser(clusterID);
    return hasNamespaceRoles(roles, dbRoles, namespace);
  }

  return true;
}

// returns true iff the current user satisfies the privileges for the set
// as specified in the arguments.
export function canAccessSet(clusterID, namespace, set, dbRoles = [],
                    clusterProps = {}, isEnterprise = false, amcRoles = []) {
  if (!hasBasicAccess(clusterID, clusterProps, isEnterprise, amcRoles))
      return false;

  const props = ClusterProps[clusterID] || {};
  if (props.isSecure) {
    const roles = getRolesForUser(clusterID);
    return hasSetRoles(roles, dbRoles, namespace, set);
  }

  return true;
}

function getRolesForUser(clusterID) {
  let croles = ClusterRoles[clusterID] || [];
  let uroles = UserRoles[clusterID] || [];
  let roles = [];

  croles.forEach((cr) => {
    if (uroles.find((r) => r === cr.name))
      roles.push(cr);
  });

  return roles;
}

function getPrivileges(roles, fn) {
  const privileges = [];
  if (typeof(fn) !== 'function')
    fn = (r) => !r.namespace && !r.set;

  if (roles) {
    roles.forEach((role) => {
      role.roles.forEach((r) => {
        if (fn(r))
          privileges.push(r.privilege);
      });
    });
  }

  return privileges;
}

function getNamespacePrivileges(roles, namespace) {
  const fn = (r) => !r.set && r.namespace === namespace;
  return getPrivileges(roles, fn);
}

function getSetPrivileges(roles, namespace, set) {
  const fn = (r) => r.namespace === namespace && r.set === set;
  return getPrivileges(roles, fn);
}

// hasRoles returns true iff the user has the required roles for the
// view type
function hasRoles(roles, required) {
  const p = getPrivileges(roles);
  return hasAny(p, required);
}

function hasNamespaceRoles(roles, required, namespace) {
  let p = getPrivileges(roles);
  if (hasAny(p, required))
    return true;

  p = getNamespacePrivileges(roles, namespace);
  return hasAny(p, required);
}

function hasSetRoles(roles, required, namespace, set) {
  if (hasNamespaceRoles(roles, required, namespace))
    return true;

  const p = getSetPrivileges(roles, namespace, set);
  return hasAny(p, required);
}

// returns true iff a has any of the values in b
function hasAny(a = [], b = []) {
  if (b.length === 0)
    return true;

  let has = false;
  b.forEach((bb) => {
    const i = a.findIndex((aa) => bb === aa);
    if (i !== -1)
      has = true;
  });
  return has;
}

// returns true iff a is a subset of b
function isArrSubset(a, b = []) {
  let subset = true;
  a.forEach((aa) => {
    const i = b.findIndex((bb) => bb === aa);
    if (i === -1)
      subset = false;
  });
  return subset;
}

// checks if all properties of a have equivalent values in b
function isPropsSubset(a, b) {
  a = a || {};
  b = b || {};

  for (let k in a) {
    let aa = a[k];
    let bb = b[k];
    if (bb === undefined) {
      if (aa)
        return false;
    } else if (bb !== aa) {
      return false;
    }
  }
  return true;
}

