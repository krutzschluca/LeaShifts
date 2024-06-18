import { getEmployees, addEmployee, updateEmployee, deleteEmployee } from './services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  const employeeList = document.getElementById('employeeList');
  const errorMessage = document.getElementById('errorMessage');
  const addEmployeeButton = document.getElementById('addEmployeeButton');
  const employeeModal = new bootstrap.Modal(document.getElementById('employeeModal'));
  const employeeForm = document.getElementById('employeeForm');
  const employeeIdInput = document.getElementById('employeeId');
  const nameInput = document.getElementById('name');

  const renderEmployees = (employees) => {
    employeeList.innerHTML = '';
    employees.forEach(employee => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${employee.id}</td>
        <td>${employee.name}</td>
        <td>
          <button class="btn btn-info btn-sm" data-employee-id="${employee.id}" data-action="edit">Edit</button>
          <button class="btn btn-danger btn-sm" data-employee-id="${employee.id}" data-action="delete">Delete</button>
        </td>
      `;
      employeeList.appendChild(row);
    });

    document.querySelectorAll('button[data-action="edit"]').forEach(button => {
      button.addEventListener('click', (event) => {
        const employeeId = event.target.getAttribute('data-employee-id');
        const employee = employees.find(employee => employee.id == employeeId);
        if (employee) {
          employeeIdInput.value = employee.id;
          nameInput.value = employee.name;
          employeeModal.show();
        }
      });
    });

    document.querySelectorAll('button[data-action="delete"]').forEach(button => {
      button.addEventListener('click', async (event) => {
        const employeeId = event.target.getAttribute('data-employee-id');
        try {
          await deleteEmployee(employeeId);
          fetchAndRenderEmployees();
        } catch (error) {
          console.error('Error deleting employee:', error);
          errorMessage.textContent = 'Failed to delete employee. Maybe he has shifts assigned.';
        }
      });
    });
  };

  const fetchAndRenderEmployees = () => {
    getEmployees().then(employees => {
      renderEmployees(employees);
    }).catch(error => {
      console.error('Error fetching employees:', error);
      errorMessage.textContent = 'Failed to load employees. Please try again later.';
    });
  };

  addEmployeeButton.addEventListener('click', () => {
    employeeForm.reset();
    employeeIdInput.value = '';
    employeeModal.show();
  });

  employeeForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    const employeeData = {
      name: nameInput.value
    };

    try {
      if (employeeIdInput.value) {
        await updateEmployee(employeeIdInput.value, employeeData);
      } else {
        await addEmployee(employeeData);
      }
      employeeModal.hide();
      fetchAndRenderEmployees();
    } catch (error) {
      console.error('Error saving employee:', error);
      errorMessage.textContent = 'Failed to save employee. Please try again later.';
    }
  });

  fetchAndRenderEmployees();
});
