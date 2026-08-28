package api

const OpenAPISpec = `openapi: 3.0.3
info:
  title: Hospital Department Account Portal API
  version: 1.0.0
  description: Deterministic API for account, department, and duty management.
servers:
  - url: http://localhost:8080
tags:
  - name: health
    description: Process health
  - name: accounts
    description: Clinical identity management
  - name: departments
    description: Department directory management
  - name: shifts
    description: Duty schedule management
paths:
  /healthz:
    get:
      tags:
        - health
      summary: Health check
      operationId: getHealthcheck
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/accounts:
    get:
      tags:
        - accounts
      summary: List accounts
      operationId: getListaccounts
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/accounts:
    post:
      tags:
        - accounts
      summary: Create account
      operationId: postCreateaccount
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/accounts/{id}:
    get:
      tags:
        - accounts
      summary: Get account
      operationId: getGetaccount
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/accounts/{id}/activate:
    post:
      tags:
        - accounts
      summary: Activate account
      operationId: postActivateaccount
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/accounts/{id}/suspend:
    post:
      tags:
        - accounts
      summary: Suspend account
      operationId: postSuspendaccount
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/departments:
    get:
      tags:
        - departments
      summary: List departments
      operationId: getListdepartments
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/departments:
    post:
      tags:
        - departments
      summary: Create department
      operationId: postCreatedepartment
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/departments/{id}:
    get:
      tags:
        - departments
      summary: Get department
      operationId: getGetdepartment
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/departments/{id}:
    put:
      tags:
        - departments
      summary: Update department
      operationId: putUpdatedepartment
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/departments/{id}/activate:
    post:
      tags:
        - departments
      summary: Activate department
      operationId: postActivatedepartment
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/shifts:
    get:
      tags:
        - shifts
      summary: List shifts
      operationId: getListshifts
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/shifts:
    post:
      tags:
        - shifts
      summary: Create shift
      operationId: postCreateshift
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
  /api/shifts/{id}/publish:
    post:
      tags:
        - shifts
      summary: Publish shift
      operationId: postPublishshift
      parameters:
        - $ref: '#/components/parameters/RequestID'
      responses:
        '200':
          description: Successful operation
          headers:
            X-Request-ID:
              schema:
                type: string
          content:
            application/json:
              schema:
                type: object
        '400':
          $ref: '#/components/responses/BadRequest'
        '403':
          $ref: '#/components/responses/Forbidden'
        '404':
          $ref: '#/components/responses/NotFound'
        '409':
          $ref: '#/components/responses/Conflict'
        '422':
          $ref: '#/components/responses/ValidationFailed'
        '500':
          $ref: '#/components/responses/InternalError'
components:
  parameters:
    RequestID:
      name: X-Request-ID
      in: header
      required: false
      schema:
        type: string
        maxLength: 128
  responses:
    BadRequest:
      description: Malformed request
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: bad_request
              message: Malformed request
              request_id: req-example
    Forbidden:
      description: Permission denied
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: permission_denied
              message: Permission denied
              request_id: req-example
    NotFound:
      description: Resource not found
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: resource_not_found
              message: Resource not found
              request_id: req-example
    Conflict:
      description: State conflict
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: state_conflict
              message: State conflict
              request_id: req-example
    ValidationFailed:
      description: Validation failed
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: validation_failed
              message: Validation failed
              request_id: req-example
    InternalError:
      description: Internal server error
      content:
        application/json:
          schema:
            $ref: '#/components/schemas/ErrorBody'
          example:
            error:
              code: internal_error
              message: Internal server error
              request_id: req-example
  schemas:
    UserAccount:
      type: object
      additionalProperties: false
      properties:
        id:
          type: string
        display_name:
          type: string
        employee_number:
          type: string
        email:
          type: string
        phone:
          type: string
        role:
          type: string
          enum: [doctor, nurse, administrator]
        department_id:
          type: string
        status:
          type: string
          description: Lifecycle state
        version:
          type: integer
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time
      required:
        - id
        - display_name
        - employee_number
        - email
        - phone
    Department:
      type: object
      additionalProperties: false
      properties:
        id:
          type: string
        code:
          type: string
        name:
          type: string
        description:
          type: string
        location:
          type: string
        phone:
          type: string
        email:
          type: string
        head_account_id:
          type: string
        status:
          type: string
          description: Lifecycle state
        services:
          type: array
        version:
          type: integer
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time
      required:
        - id
        - code
        - name
        - description
        - location
    DutyShift:
      type: object
      additionalProperties: false
      properties:
        id:
          type: string
        department_id:
          type: string
        account_id:
          type: string
        title:
          type: string
        start_at:
          type: string
          format: date-time
        end_at:
          type: string
          format: date-time
        status:
          type: string
          description: Lifecycle state
        notes:
          type: string
        version:
          type: integer
        created_at:
          type: string
          format: date-time
        updated_at:
          type: string
          format: date-time
      required:
        - id
        - department_id
        - account_id
        - title
        - start_at
    AuditRecord:
      type: object
      additionalProperties: false
      properties:
        id:
          type: string
        actor_id:
          type: string
        action:
          type: string
        subject_type:
          type: string
        subject_id:
          type: string
        outcome:
          type: string
        fields:
          type: object
        occurred_at:
          type: string
          format: date-time
      required:
        - id
        - actor_id
        - action
        - subject_type
        - subject_id
    ErrorBody:
      type: object
      additionalProperties: false
      properties:
        error:
          type: object
      required:
        - error`
