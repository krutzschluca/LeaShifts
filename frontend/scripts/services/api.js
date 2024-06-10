const API_BASE_URL = 'https://leashifts.onrender.com';

// Generic GET request function
const fetchData = async (url) => {
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Fetch error:', error);
    throw error;
  }
};

// Generic POST request function
const postData = async (url, body) => {
  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Post error:', error);
    throw error;
  }
};

// Generic PUT request function
const putData = async (url, body) => {
  try {
    const response = await fetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(body)
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Put error:', error);
    throw error;
  }
};

// Generic DELETE request function
const deleteData = async (url) => {
  try {
    const response = await fetch(url, {
      method: 'DELETE'
    });
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    // If there is no content in the response, return an empty object or appropriate message
    if (response.status === 204) {
      return {};
    }
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Delete error:', error);
    throw error;
  }
};

// Employee API functions
export const getEmployees = () => fetchData(`${API_BASE_URL}/employees`);
export const getEmployeeById = (id) => fetchData(`${API_BASE_URL}/employees/${id}`);
export const addEmployee = (employee) => postData(`${API_BASE_URL}/employees`, employee);
export const updateEmployee = (id, employee) => putData(`${API_BASE_URL}/employees/${id}`, employee);
export const deleteEmployee = (id) => deleteData(`${API_BASE_URL}/employees/${id}`);

// Store API functions
export const getStores = () => fetchData(`${API_BASE_URL}/stores`);
export const getStoreById = (id) => fetchData(`${API_BASE_URL}/stores/${id}`);
export const addStore = (store) => postData(`${API_BASE_URL}/stores`, store);
export const updateStore = (id, store) => putData(`${API_BASE_URL}/stores/${id}`, store);
export const deleteStore = (id) => deleteData(`${API_BASE_URL}/stores/${id}`);

// Shift API functions
export const getShifts = () => fetchData(`${API_BASE_URL}/shifts`);
export const getShiftById = (id) => fetchData(`${API_BASE_URL}/shifts/${id}`);
export const addShift = (shift) => postData(`${API_BASE_URL}/shifts`, shift);
export const updateShift = (id, shift) => putData(`${API_BASE_URL}/shifts/${id}`, shift);
export const deleteShift = (id) => deleteData(`${API_BASE_URL}/shifts/${id}`);
