import { addShift, getEmployees, getStores } from './services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  const shiftForm = document.getElementById('shiftForm');
  const storeIdSelect = document.getElementById('storeId');
  const employeeIdSelect = document.getElementById('employeeId');
  const startingTimeSelect = document.getElementById('startingTime');
  const endingTimeSelect = document.getElementById('endingTime');
  const successMessage = document.getElementById('successMessage');
  const errorMessage = document.getElementById('errorMessage');

  const populateTimeDropdown = (selectElement, openingTime, closingTime) => {
    selectElement.innerHTML = '';
    let start = new Date(`1970-01-01T${openingTime}Z`);
    let end = new Date(`1970-01-01T${closingTime}Z`);

    while (start < end) {
      const timeString = start.toISOString().substring(11, 16);
      const option = document.createElement('option');
      option.value = timeString;
      option.textContent = timeString;
      selectElement.appendChild(option);
      start.setMinutes(start.getMinutes() + 15);
    }
  };

  getStores().then(stores => {
    stores.forEach(store => {
      const option = document.createElement('option');
      option.value = store.id;
      option.textContent = store.address;
      storeIdSelect.appendChild(option);
    });
  }).catch(error => {
    console.error('Error fetching stores:', error);
    errorMessage.textContent = 'Failed to load stores. Please try again later.';
  });

  getEmployees().then(employees => {
    employees.forEach(employee => {
      const option = document.createElement('option');
      option.value = employee.id;
      option.textContent = employee.name;
      employeeIdSelect.appendChild(option);
    });
  }).catch(error => {
    console.error('Error fetching employees:', error);
    errorMessage.textContent = 'Failed to load employees. Please try again later.';
  });

  storeIdSelect.addEventListener('change', () => {
    const selectedStore = storeIdSelect.value;
    getStores().then(stores => {
      const store = stores.find(store => store.id == selectedStore);
      if (store) {
        populateTimeDropdown(startingTimeSelect, store.opening_time, store.closing_time);
        populateTimeDropdown(endingTimeSelect, store.opening_time, store.closing_time);
      }
    }).catch(error => {
      console.error('Error fetching stores:', error);
      errorMessage.textContent = 'Failed to load store details. Please try again later.';
    });
  });

  shiftForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    
    const storeId = document.getElementById('storeId').value;
    const employeeId = document.getElementById('employeeId').value;
    const date = document.getElementById('date').value;
    const startingTime = document.getElementById('startingTime').value;
    const endingTime = document.getElementById('endingTime').value;

    const shiftData = {
      store_id: parseInt(storeId),
      employee_id: parseInt(employeeId),
      date: new Date(date).toISOString(),
      starting_time: startingTime,
      ending_time: endingTime
    };

    try {
      await addShift(shiftData);
      successMessage.textContent = 'Shift added successfully!';
      errorMessage.textContent = '';
      shiftForm.reset();
    } catch (error) {
      console.error('Error adding shift:', error);
      successMessage.textContent = '';
      errorMessage.textContent = 'Failed to add shift. Please try again later.';
    }
  });
});
