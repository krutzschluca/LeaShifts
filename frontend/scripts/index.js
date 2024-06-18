import { getShifts, getStores, getEmployees, deleteShift } from './services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  const shiftList = document.getElementById('shiftList');
  const errorMessage = document.getElementById('errorMessage');
  const prevWeekButton = document.getElementById('prevWeek');
  const nextWeekButton = document.getElementById('nextWeek');
  const storeSelect = document.getElementById('storeSelect');
  const currentWeekDisplay = document.getElementById('currentWeek');
  const deleteModal = new bootstrap.Modal(document.getElementById('deleteModal'));
  const confirmDeleteButton = document.getElementById('confirmDelete');
  let shiftIdToDelete = null;

  let currentDate = new Date();
  let employeeMap = {};

  const updateWeekDisplay = () => {
    const startOfWeek = new Date(currentDate.setDate(currentDate.getDate() - currentDate.getDay()));
    const endOfWeek = new Date(currentDate.setDate(startOfWeek.getDate() + 6));
    currentWeekDisplay.textContent = `Week of ${startOfWeek.toLocaleDateString()} - ${endOfWeek.toLocaleDateString()}`;
  };

  const renderShifts = (shifts) => {
    shiftList.innerHTML = '';
    shifts.forEach(shift => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${shift.id}</td>
        <td>${employeeMap[shift.employee_id]}</td>
        <td>${new Date(shift.date).toLocaleDateString()}</td>
        <td>${shift.starting_time}</td>
        <td>${shift.ending_time}</td>
        <td>
          <button class="btn btn-danger btn-sm" data-shift-id="${shift.id}" data-toggle="modal" data-target="#deleteModal">Delete</button>
        </td>
      `;
      shiftList.appendChild(row);
    });

    if (shifts.length === 0) {
      shiftList.innerHTML = '<tr><td colspan="6" class="text-center text-warning">No shifts scheduled for this week.</td></tr>';
    }

    // Add event listeners for delete buttons
    document.querySelectorAll('button[data-shift-id]').forEach(button => {
      button.addEventListener('click', (event) => {
        shiftIdToDelete = event.target.getAttribute('data-shift-id');
        $('#deleteModal').modal('show');
      });
    });
  };

  const fetchAndRenderShifts = () => {
    const selectedStoreId = storeSelect.value;
    if (!selectedStoreId) {
      shiftList.innerHTML = '<tr><td colspan="6" class="text-center text-warning">Please select a store to view shifts.</td></tr>';
      return;
    }

    getShifts().then(shifts => {
      const startOfWeek = new Date(currentDate.setDate(currentDate.getDate() - currentDate.getDay()));
      const endOfWeek = new Date(currentDate.setDate(startOfWeek.getDate() + 6));

      const isCurrentWeek = (date) => {
        const shiftDate = new Date(date);
        return shiftDate >= startOfWeek && shiftDate <= endOfWeek;
      };

      const currentWeekShifts = shifts.filter(shift => isCurrentWeek(shift.date) && shift.store_id == selectedStoreId);
      renderShifts(currentWeekShifts);
      updateWeekDisplay();
    }).catch(error => {
      console.error('Error fetching shifts:', error);
      errorMessage.textContent = 'Failed to load shifts. Please try again later.';
    });
  };

  const fetchStoresAndEmployees = () => {
    Promise.all([getStores(), getEmployees()])
      .then(([stores, employees]) => {
        // Populate store dropdown
        stores.forEach(store => {
          const option = document.createElement('option');
          option.value = store.id;
          option.textContent = store.address;
          storeSelect.appendChild(option);
        });

        employeeMap = employees.reduce((map, employee) => {
          map[employee.id] = employee.name;
          return map;
        }, {});

        fetchAndRenderShifts();
      })
      .catch(error => {
        console.error('Error fetching stores or employees:', error);
        errorMessage.textContent = 'Failed to load stores or employees. Please try again later.';
      });
  };

  storeSelect.addEventListener('change', fetchAndRenderShifts);
  prevWeekButton.addEventListener('click', () => {
    currentDate.setDate(currentDate.getDate() - 7);
    fetchAndRenderShifts();
  });

  nextWeekButton.addEventListener('click', () => {
    currentDate.setDate(currentDate.getDate() + 7);
    fetchAndRenderShifts();
  });

  confirmDeleteButton.addEventListener('click', () => {
    if (shiftIdToDelete) {
      deleteShift(shiftIdToDelete).then(() => {
        $('#deleteModal').modal('hide');
        fetchAndRenderShifts();
      }).catch(error => {
        console.error('Error deleting shift:', error);
        errorMessage.textContent = 'Failed to delete shift. Please try again later.';
      });
    }
  });

  fetchStoresAndEmployees();
});
